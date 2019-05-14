Multithreaded City Trafic Simulator
===================================

This is an implementation of a city traffic simulator in which cars and semaphores are independent entities, semaphores have to be respected by the vehicles. 

Each entity will be used by a separate thread so everything requires to be syncronized.


Features
--------

- The city map is static 
- Any number (up to 12) of vehicles can be added.
- All vehicles have a random start and end point (u-turns are prohibited).
- Vehicles have random speeds.
- Car proximity detection (vehicles slow down if another is close by). 
- Multithreaded implementation (vehicles and semaphores are in separated threads).
- All entities are on the same layout. 
- Vehicle routes are displayed at the end.
- Vehicle speeds are displayed.

Dependency Installation
-----------------------

### On __Windows__ based systems

1. Install mingw-w64 from [Mingw-builds](http://mingw-w64.org/doku.php/download/mingw-builds)
    * Version: latest (at time of writing 6.3.0)
    * Architecture: x86_64
    * Threads: win32
    * Exception: seh
    * Build revision: 1
    * Destination Folder: Select a folder that your Windows user owns
2. Install SDL2 http://libsdl.org/download-2.0.php
    * Extract the SDL2 folder from the archive using a tool like [7zip](http://7-zip.org)
    * Inside the folder, copy the `i686-w64-mingw32` and/or `x86_64-w64-mingw32` depending on the architecture you chose into your mingw-w64 folder e.g. `C:\Program Files\mingw-w64\x86_64-6.3.0-win32-seh-rt_v5-rev1\mingw64`
3. Setup Path environment variable
    * Put your mingw-w64 binaries location into your system Path environment variable. e.g. `C:\Program Files\mingw-w64\x86_64-6.3.0-win32-seh-rt_v5-rev1\mingw64\bin` and `C:\Program Files\mingw-w64\x86_64-6.3.0-win32-seh-rt_v5-rev1\mingw64\x86_64-w64-mingw32\bin`
4. Open up a terminal such as `Git Bash` and run `go get -v github.com/veandco/go-sdl2/sdl`.
5. (Optional) You can repeat __Step 2__ for [SDL_image](https://www.libsdl.org/projects/SDL_image), [SDL_mixer](https://www.libsdl.org/projects/SDL_mixer), [SDL_ttf](https://www.libsdl.org/projects/SDL_ttf)
    * NOTE: pre-build the libraries for faster compilation by running `go install github.com/veandco/go-sdl2/{sdl,img,mix,ttf}`

* Or you can install SDL2 via [Msys2](https://msys2.github.io) like so:
`pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-SDL2{,_image,_mixer,_ttf,_gfx}`

### On __Linux__ based systems

On __Ubuntu 14.04 and above__, type:\
`apt install libsdl2{,-image,-mixer,-ttf,-gfx}-dev`

On __Fedora 25 and above__, type:\
`yum install SDL2{,_image,_mixer,_ttf,_gfx}-devel`

On __Arch Linux__, type:\
`pacman -S sdl2{,_image,_mixer,_ttf,_gfx}`

On __Gentoo__, type:\
`emerge -av libsdl2 sdl2-{image,mixer,ttf,gfx}`

### On __MacOs__ based systems

install SDL2 via [Homebrew](http://brew.sh) like so:\
`brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config`

Compilation
-----------

__*Clone repository into the $GOPATH/src/ folder*__

### On __Linux__ or __MacOs__

```bash
# To compile the go proyect
make

# To remove all files (including executable) created after compilation.
make clean
```

### On __Windows__

```bash
# To get dependencies
go get -v

# To build the proyect
go build
```

How to run
----------

### On __Linux__ or __MacOs__

```bash
# -ncars=num where num = number of cars (from 1 up to 12)
./CityTraffic [-ncars=4]
```
### On __Windows__

```bash
# -ncars=num where num = number of cars (from 1 up to 12)
# replace foldername with the name of the folder in which the go file was built
city-traffic.exe [-ncars=4]
```

License
-------

The project is licensed under the MIT license.
