# Go Wiper
![image](https://user-images.githubusercontent.com/43169346/141783010-3463f82b-c1aa-4bb0-b662-72d4a9d09efa.png)

You can use this tool like a library or a ready program. If you thought in some times about safely data erasing, you have a great open-source tool that can safely delete your secure data without any recovery possibilities.

You can add own implementation of wiping rule from wipe.Rule interface. 
But, there is already a ready rules like a Peter Gutmann (35 passes), or US Department of Defense DoD 5220.22-M (3 passes), etc...

## Usage
1. Clone the repository `git clone https://github.com/0x9ef/go-wiper`
2. Then you must build a program `golang build gowiper.go`
3. After that you can run the program and get some helpful information `./gowiper --help` 

## License
https://github.com/0x9ef/go-wiper/blob/master/LICENSE
