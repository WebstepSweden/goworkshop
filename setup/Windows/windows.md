## Windows

1. Download GoLand from [https://www.jetbrains.com/go/](https://www.jetbrains.com/go/)
![Download GoLand](download_goland.jpg)

1. Once downloaded, install it
![Install GoLand](install_goland.jpg)

1. Create folders *C:\go* and *C:\go\src*
![Create folders](create_folders.jpg)

1. Clone the workshop repo from [https://github.com/WebstepSweden/goworkshop.git](https://www.jetbrains.com/go/) into the *C:\go\src* folder
![Clone workshop](clone_workshop.jpg)

1. Open the *C:\go* folder with GoLand
![Open workshop](open_workshop.jpg)

1. Go to *File* > *Settings* > *Go* > *GOROOT* and choose *Add SDK...* > *Download...*
![Choose SDK](download_sdk_1.jpg)

1. Select the latest version and click *Ok*
![Download SDK](download_sdk_2.jpg)

1. Add the bin folder inside the downloaded sdk to your PATH environemt variable
![Add bin to PATH](add_go_to_path.jpg)

1. Go to *File* > *Settings* > *Go* > *GOPATH* and add the *C:\go* folder to the *Global GOPATH*
![Add to GOPATH](add_gopath.jpg)

1. Open the *try_me_first_test.go* file and run the test by clicking the green arrow beside *func TestIfThisWorks...*. If everything works fine, you should see a confirmation in the run window
![Run first test](run_first_test.jpg)
