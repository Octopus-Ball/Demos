"""
逐个画面遍历视频，找到合适的片段
"""
from tools import if_black, rgb2hsb

class VideoInspector:
    def __init__(self, video, boundary_marker, back_color):
        self.video = video
        self.boundary_marker = boundary_marker      # 边界检测异物的范围
        self.back_color = back_color                # 背景颜色信息
        self.ok_time_li = []                        # 可保留的画面的时间点的列表
        self.cut_time_round_li = []                 # 视频剪切时间范围
        # 可调整参数
        self.interval = 0.1                         # 取样间隔(0.1秒)
        self.side = 2                               # 排除四边各2px不予计算
        self.h_tolerance = 100                      # 检测色相是否相等的容差为30
        self.s_tolerance = 10                       # 检测饱和度是否相等的容差为20
        self.impurities_tolerance = 500             # 杂质容差(检测到手的像素点超过100个，则视为需要剪掉的片段)

        self.inspect_video()
        self.get_cut_time_range()


    def inspect_video(self):
        # 遍历视频，找到可保留画面的时间点
        t = 0
        while t <= self.video.end:
            if self.inspect_frame(t):
                self.ok_time_li.append(t)
                print("该时间点可保留", t)
            t = round((t + self.interval), 1)

    def inspect_frame(self, t):
        # 检查时间t所在一帧画面是否为可保留画面
        impurities_sum = 0      # 统计的杂质点的数量
        RGB = self.frame_rgb(t)
        for rgb in RGB:
            if not if_black(rgb):
                h, s, _ = rgb2hsb(rgb)
                if not self.inspect_dot(h, s):
                    impurities_sum += 1
                    if impurities_sum >= self.impurities_tolerance:
                        return False
        return True

    def inspect_dot(self, h, s):
        # 检查某点的h和s值是否与背景色相等
        if abs(self.back_color["h"] - h) > self.h_tolerance:
            return False
        else:
            if s <= (self.back_color["s"] - self.s_tolerance):
                return False
        return True

    def frame_rgb(self, t):
        # 某时间点的帧画面里检测区域内的rgb信息的生成器
        all_rgb = self.video.get_frame(t)
        top_range = (self.side, self.boundary_marker["top"])
        bottom_range = (self.boundary_marker["bottom"], self.video.h-self.side)
        left_range = (self.side, self.boundary_marker["left"])
        right_range = (self.boundary_marker["right"], self.video.w-self.side)

        for line in range(top_range[0], top_range[1]):
            for rgb in all_rgb[line]:
                yield rgb
        for line in range(bottom_range[0], bottom_range[1]):
            for rgb in all_rgb[line]:
                yield rgb
        for line in range(top_range[1], bottom_range[0]):
            for row in range(left_range[0], left_range[1]):
                rgb = all_rgb[line][row]
                yield rgb
            for row in range(right_range[0], right_range[1]):
                rgb = all_rgb[line][row]
                yield rgb

    def get_cut_time_range(self):
        # 确定视频剪切时间区间
        li = self.ok_time_li
        start_p = end_p = 0
        tmp_p = end_p + 1
        while tmp_p < len(li):
            if li[tmp_p] == round(li[end_p] + 0.1, 1):
                end_p += 1
                tmp_p += 1
            else:
                start_time = li[start_p]
                end_time = li[end_p]
                self.cut_time_round_li.append([start_time, end_time])
                start_p = end_p = tmp_p
                tmp_p += 1
        start_time = li[start_p]
        end_time = li[len(li)-1]
        self.cut_time_round_li.append((start_time, end_time))