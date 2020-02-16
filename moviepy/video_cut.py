"""
根据剪辑时间区间剪辑视频
"""
from moviepy.editor import concatenate_videoclips

class VidoeCut:
    def __init__(self, video, cut_time_range_li, to_path):
        self.video = video
        self.cut_time_range_li = cut_time_range_li
        self.to_path = to_path

        self.cut_video()

    def cut_video(self):
        clip_arr = []
        for time_range in self.cut_time_range_li:
            start_time, end_time = time_range
            start_time += 0.1
            end_time -= 0.1
            clip = self.video.subclip((start_time), (end_time))
            clip_arr.append(clip)
        finalclip = concatenate_videoclips(clip_arr)
        finalclip.write_videofile(self.to_path)