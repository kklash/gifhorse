# gifhorse

Quickly convert a .mp4 or other video file into a gif, with powerful but simple command line options. `gifhorse` is just a wrapper for `ffmpeg`, so you'll have to install that first.

## Install

You will need to install `ffmpeg` if you don't already have it on your system. Check by running `ffmpeg -version`.

On Debian/Ubuntu, a simple `apt-get` will usually suffice:
```
$ sudo apt-get install ffmpeg
```

On Mac OS, you can use homebrew:
```
$ brew install ffmpeg
```

Then to install `gifhorse`, download the signed binary tarball for your OS from the Releases page. The releases come with signed SHA256 hashes, and [the signing PGP key is available via Keybase](https://keybase.io/kklashinsky/pgp_keys.asc?fingerprint=4104f024165e528a71ac784bcfbf781763b4d7af).

```
$ keybase id kklashinsky
▶ INFO Identifying kklashinsky
✔ public key fingerprint: 4104 F024 165E 528A 71AC 784B CFBF 7817 63B4 D7AF
✔ "exodus_konnor" on reddit: https://www.reddit.com/r/KeybaseProofs/comments/73ce6k/my_keybase_proof_redditexodus_konnor/
✔ "kklash" on github: https://gist.github.com/6647e3721da17b3ccb99f83dffa27e39
```

If you don't have GPG, [you can verify the SHA256 release hashes here](https://keybase.io/verify).

Example installing to `/usr/local/bin`:
```
$ tar -xf gifhorse-darwin_dev_amd64.tar.gz
$ cd gifhorse-darwin_dev_amd64/
$ sudo cp gifhorse /usr/local/bin
$ sudo chmod 755 /usr/local/bin/gifhorse # fix permissions
```

## From Source

You'll need a `go` compiler:
```sh
$ sudo apt-get install golang # linux
$ brew install golang # mac
```

From there, it's simple to install:
```sh
$ go get -v github.com/kklash/gifhorse
```

## Usage
```
$ gifhorse  --help
Usage of gifhorse:
  -debug
    	debug mode, print ffmpeg output
  -duration float
    	duration of the video (in seconds) from -offset to convert from
  -force
    	automatically overwrite output files
  -framerate int
    	output gif framerate (default 15)
  -height int
    	output image frame height (in pixels). incompatible with -resize
  -in string
    	input video file name (REQUIRED)
  -offset float
    	offset time (in seconds) from start of video to convert from
  -out string
    	output .gif file name (must end with .gif) (default "out.gif")
  -resize float
    	image resizing multiplier. for example, '-resize 0.5' means the output frame size is half of the input size.
  -width int
    	output image frame width (in pixels). incompatible with -resize
```
## Examples

##### The Basics
```sh
$ gifhorse -in video.mp4 -out video.gif
Converted to gif, size 42.29 MB
video.mp4 -> video.gif
```

```sh
$ gifhorse -in /home/user/Desktop/my_vid.mp4
Converted to gif, size 42.29 MB
/home/user/Desktop/my_vid.mp4 -> out.gif # Outputs to `out.gif` in CWD by default
```

##### Optimization
The `-resize` multiplier option allows you to shrink or enlarge a gif during conversion while keeping the original aspect ratio.
```sh
# Using -resize, you can shrink the output frame size for smaller
# file sizes, but keep the aspect ratio of the original video.
$ gifhorse -in video.mp4 -resize 0.333 -out resized_video.gif
```

If you want more precise control, you can use the `-width` and/or `-height` flags.
```sh
$ gifhorse -in video.mp4 -height 720 -width 1280 -out video_1280x720.gif
```

If you only supply one of `-width` or `-height`, `ffmpeg` will automagically keep the original video's aspect ratio in check.
```sh
$ gifhorse -in video_1920x1080.mp4 -height 720 -out video_1280x720.gif
```

You can even perform resizing operations on gif input files!
```sh
$ gifhorse -in big.gif -resize 0.2 -out small.gif
```

Adjust the output gif's framerate manually with `-framerate`.
```sh
$ gifhorse -in high_quality_video.mp4 -framerate 30 -out high_quality_gif.gif
$ gifhorse -in high_quality_video.mp4 -framerate 5 -out low_quality_gif.gif
```

##### Timing
Use the `-offset` flag to set the timestamp (in seconds) in the video where you want to start the conversion from.
```sh
$ gifhorse -in long_video.mp4 -offset 17.5 -out shorter_video.gif
```

Use the `-duration` flag to choose the length of time from the video you want to convert.
```sh
$ gifhorse -in long_video.mp4 -duration 4.2 -out first_4_seconds.gif
```

Combine `-duration` and `-offset` to splice out precise segments of a video.
```sh
# This would convert the video between 0:30 -> 0:40
$ gifhorse -in something_exciting_happens_at_30_second_mark.mp4 \
  -offset 30 \
  -duration 10 \
  -out exciting_things_happening.gif
```
