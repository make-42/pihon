ffmpeg -i pihon_demo.mp4 -filter_complex "[0:v]reverse,fifo[r];[0:v][r] concat=n=2:v=1 [v]" -map "[v]" -crf 2 -preset veryslow outputLooping.mp4