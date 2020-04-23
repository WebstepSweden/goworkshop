## Mac

1. Download GoLand from [https://www.jetbrains.com/go/](https://www.jetbrains.com/go/)
![Download GoLand](download_goland.png)

1. Once downloaded, install it
![Install GoLand](install_goland.png)

1. Create folders *~/go* and *~/go/src*
![Create folders](create_folders.png)

1. Clone the workshop repo from [https://github.com/WebstepSweden/goworkshop.git](https://github.com/WebstepSweden/goworkshop.git) into the *~/go/src* folder
![Clone workshop](clone_workshop.png)

1. Open the *~/go* folder with GoLand
![Open workshop](open_workshop.png)

1. Go to *GoLand* > *Preferences* > *Go* > *GOROOT* and choose *Add SDK...* > *Download...*
![Choose SDK](download_sdk_1.png)

1. Select the latest version and click *Ok* (don't install it under *~/go* though, choose a different path)
![Download SDK](download_sdk_2.png)

1. Add the bin folder inside the downloaded sdk to your PATH environemt variable in *.bash_profile* or *.zshrc* file on your home directory
![Add bin to PATH](add_go_to_path.png)

1. Back in GoLand, go to *GoLand* > *Preferences* > *Go* > *GOPATH* and add the *~/go* folder to the *Global GOPATH*
![Add to GOPATH](add_gopath.png)

1. Open the *try_me_first_test.go* file and run the test by clicking the green arrow beside *func TestIfThisWorks...*. If everything works fine, you should see a confirmation in the run window
![Run first test](run_first_test.png)
