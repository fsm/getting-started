<a href="https://github.com/fsm"><p align="center"><img src="https://user-images.githubusercontent.com/2105067/35464215-a014d512-02a9-11e8-8913-63a066f6064e.png" alt="FSM" width="350px" align="center;"/></p></a>

# Getting Started

This package contains an example project / scaffolding for [FSM](https://github.com/fsm/fsm).

## Setting Up

The first thing you're going to want to do is clone down the repository to the location you would like your project to exist.  This should be somewhere in your $GOPATH.

After you've cloned it down, you're going to need to update the package on [this line](https://github.com/fsm/getting-started/blob/master/main.go#L7) that references a nested package within this repo.  It will need to be updated to reflect the location where you cloned the repository.

## Building

After your project is set up as specified in the above section, you'll need to install dependencies.

Assuming you have [dep](https://github.com/golang/dep) installed, you can run the following command in the root of the repository:

```sh
dep ensure
```

Now you can build and run the project.

```sh
go build -o bot && ./bot
```

## License

[MIT](LICENSE.md)