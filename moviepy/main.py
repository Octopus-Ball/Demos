from moviepy.editor import VideoFileClip, concatenate_videoclips, vfx
from boundary_marker import BoundaryMarker
from video_inspector import VideoInspector
from video_cut import VidoeCut

if __name__ == "__main__":
    from_path = "../video/from.mp4"
    to_path = "../video/to.mp4"

    video = VideoFileClip(from_path)
    video = video.fx(vfx.rotate, 90)    # 旋转90度,(横屏变竖屏)

    bm = BoundaryMarker(video)
    vi = VideoInspector(video, bm.boundary_marker, bm.back_color)
    cut_time_li = vi.cut_time_round_li
    print(cut_time_li)
    # vc = VidoeCut(video, vi.cut_time_round_li, to_path)