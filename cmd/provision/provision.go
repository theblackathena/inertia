package provisioncmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/ubclaunchpad/inertia/cfg"
	"github.com/ubclaunchpad/inertia/client"
	inertiacmd "github.com/ubclaunchpad/inertia/cmd/cmd"
	"github.com/ubclaunchpad/inertia/cmd/inpututil"
	"github.com/ubclaunchpad/inertia/cmd/printutil"
	"github.com/ubclaunchpad/inertia/common"
	"github.com/ubclaunchpad/inertia/local"
	"github.com/ubclaunchpad/inertia/provision"
)

// ProvisionCmd is the parent class for the 'inertia provision' subcommands
type ProvisionCmd struct {
	*cobra.Command
	config  *cfg.Config
	cfgPath string
}

const (
	flagDaemonPort = "daemon.port"
	flagPorts      = "ports"
)

// AttachProvisionCmd attaches the 'provision' subcommands to the given parent
func AttachProvisionCmd(inertia *inertiacmd.Cmd) {
	var prov = &ProvisionCmd{}
	prov.Command = &cobra.Command{
		Use:   "provision",
		Short: "Provision a new remote host to deploy your project on",
		Long:  `Provisions a new remote host set up for continuous deployment with Inertia.`,
		PersistentPreRun: func(*cobra.Command, []string) {
			// Ensure project initialized, load config
			var err error
			prov.config, prov.cfgPath, err = local.GetProjectConfigFromDisk(inertia.ConfigPath)
			if err != nil {
				printutil.Fatalf("failed to read config at '%s': %s", prov.cfgPath, err.Error())
			}
			if prov.config == nil {
				printutil.Fatalf("failed to read config at '%s'", prov.cfgPath)
			}
		},
	}
	prov.PersistentFlags().StringP(flagDaemonPort, "d", "4303", "daemon port")
	prov.PersistentFlags().StringArrayP(flagPorts, "p", []string{}, "ports your project uses")

	// add children
	prov.attachEcsCmd()

	// add to parent
	inertia.AddCommand(prov.Command)
}

func (root *ProvisionCmd) attachEcsCmd() {
	const (
		flagType        = "type"
		flagUser        = "user"
		flagFromEnv     = "from-env"
		flagFromProfile = "from-profile"
		flagProfilePath = "profile.path"
		flagProfileUser = "profile.user"
	)
	var provEC2 = &cobra.Command{
		Use:   "ec2 [name]",
		Short: "[BETA] Provision a new Amazon EC2 instance",
		Long: `[BETA] Provisions a new Amazon EC2 instance and sets it up for continuous deployment
with Inertia. 

Make sure you run this command with the '-p' flag to indicate what ports
your project uses - for example:

	inertia provision ec2 my_ec2_instance -p 8000

This ensures that your project ports are properly exposed and externally accessible.
`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var config = root.config
			if _, found := config.GetRemote(args[0]); found {
				printutil.Fatal("remote with name already exists")
			}

			// Load flags for credentials
			var fromEnv, _ = cmd.Flags().GetBool(flagFromEnv)
			var withProfile, _ = cmd.Flags().GetBool(flagFromProfile)

			// Load flags for setup configuration
			var user, _ = cmd.Flags().GetString(flagUser)
			var instanceType, _ = cmd.Flags().GetString(flagType)
			var stringProjectPorts, _ = cmd.Flags().GetStringArray(flagPorts)
			if stringProjectPorts == nil || len(stringProjectPorts) == 0 {
				fmt.Print("[WARNING] no project ports provided - this means that no ports" +
					"will be exposed on your ec2 host. Use the '--ports' flag to set" +
					"ports that you want to be accessible.")
			}

			// Create VPS instance
			var prov *provision.EC2Provisioner
			var err error
			if fromEnv {
				prov, err = provision.NewEC2ProvisionerFromEnv(user, os.Stdout)
				if err != nil {
					printutil.Fatal(err)
				}
			} else if withProfile {
				var profileUser, _ = cmd.Flags().GetString(flagProfileUser)
				var profilePath, _ = cmd.Flags().GetString(flagProfilePath)
				prov, err = provision.NewEC2ProvisionerFromProfile(
					user, profileUser, profilePath, os.Stdout)
				if err != nil {
					printutil.Fatal(err)
				}
			} else {
				keyID, key, err := inpututil.EnterEC2CredentialsWalkthrough(os.Stdin)
				if err != nil {
					printutil.Fatal(err)
				}
				prov, err = provision.NewEC2Provisioner(user, keyID, key, os.Stdout)
				if err != nil {
					printutil.Fatal(err)
				}
			}

			// Report connected user
			fmt.Printf("Executing commands as user '%s'\n", prov.GetUser())

			// Prompt for region
			println("See https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions for a list of available regions.")
			print("Please enter a region: ")
			var region string
			if _, err = fmt.Fscanln(os.Stdin, &region); err != nil {
				printutil.Fatal(err)
			}

			// List image options and prompt for input
			fmt.Printf("Loading images for region '%s'...\n", region)
			images, err := prov.ListImageOptions(region)
			if err != nil {
				printutil.Fatal(err)
			}
			image, err := inpututil.ChooseFromListWalkthrough(os.Stdin, "image", images)
			if err != nil {
				printutil.Fatal(err)
			}

			// Gather input
			fmt.Printf("Creating %s instance in %s from image %s...\n", instanceType, region, image)
			var ports = []int64{}
			for _, portString := range stringProjectPorts {
				p, err := common.ParseInt64(portString)
				if err == nil {
					ports = append(ports, p)
				} else {
					fmt.Printf("invalid port %s", portString)
				}
			}

			// Create remote instance
			var port, _ = cmd.Flags().GetString(flagDaemonPort)
			var portDaemon, _ = common.ParseInt64(port)
			remote, err := prov.CreateInstance(provision.EC2CreateInstanceOptions{
				Name:        args[0],
				ProjectName: config.Project,
				Ports:       ports,
				DaemonPort:  portDaemon,

				ImageID:      image,
				InstanceType: instanceType,
				Region:       region,
			})
			if err != nil {
				printutil.Fatal(err)
			}

			// Save new remote to configuration
			remote.Branch, err = local.GetRepoCurrentBranch()
			if err != nil {
				printutil.Fatal(err)
			}
			config.AddRemote(remote)
			config.Write(root.cfgPath)

			// Create inertia client
			inertia, found := client.NewClient(args[0], os.Getenv(local.EnvSSHPassphrase), config, os.Stdout)
			if !found {
				printutil.Fatal("vps setup did not complete properly")
			}

			// Bootstrap remote
			fmt.Printf("Initializing Inertia daemon at %s...\n", inertia.RemoteVPS.IP)
			if err = inertia.BootstrapRemote(config.Project); err != nil {
				printutil.Fatal(err)
			}

			// Save updated config
			config.Write(root.cfgPath)
		},
	}
	provEC2.Flags().StringP(flagType, "t",
		"t2.micro", "ec2 instance type to instantiate")
	provEC2.Flags().StringP(flagUser, "u",
		"ec2-user", "ec2 instance user to execute commands as")
	provEC2.Flags().Bool(flagFromEnv, false,
		"load ec2 credentials from environment - requires AWS_ACCESS_KEY_ID, AWS_ACCESS_KEY to be set")
	provEC2.Flags().Bool(flagFromProfile, false,
		"load ec2 credentials from profile")
	provEC2.Flags().String(flagProfilePath, "~/.aws/config",
		"path to aws profile configuration file")
	provEC2.Flags().String(flagProfileUser, "default",
		"user profile for aws credentials file")

	root.AddCommand(provEC2)
}
