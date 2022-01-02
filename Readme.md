# Subtitle Renamifier

![Image](https://github.com/majidalaeinia/subtitle-renamifier/blob/master/image.png?raw=true)

You might have a folder containing different episodes of a serial and its corresponding subtitles.
This repo helps you rename the subtitle files with their corresponding video name.

Suppose that you have such content in your folder. Just put the executable file on this folder and press the button.
```
├── BreakingBad.S01E01.mkv
├── BreakingBad.S01E02.mkv
├── BreakingBad.S01E03.mkv
├── donyayezirnevis.com.breakingbad.Bluray.series.S01E02.srt
├── subtitleworld.com.breakingbad.720.S01E01.srt
└── subtitleworld.com.breakingbad.720.S01E03.srt
```
The output will be:
```
├── BreakingBad.S01E01.mkv
├── BreakingBad.S01E02.mkv
├── BreakingBad.S01E03.mkv
├── BreakingBad.S01E01.srt
├── BreakingBad.S01E02.srt
└── BreakingBad.S01E03.srt
```

## TODO
- [ ] Write a descriptive readme
- [ ] Allow user to choose the video extension
- [ ] Allow user to choose the regex pattern (Season and Episode name pattern)
- [ ] Show error on a separated dialog box
- [ ] Beautify the window
- [ ] Add tests
- [ ] Accept custom path
- [ ] Handle leading zeros
- [ ] Make downloadable executable file for Windows, Linux and iOS
