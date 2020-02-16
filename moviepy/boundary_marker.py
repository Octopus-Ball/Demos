"""
通过黑色标定点确定画面边界
"""
from tools import if_black, rgb2hsb, get_most_value

class BoundaryMarker:
    def __init__(self, video):
         # 可自定义调整参数
        self.point_time = 0                 # 采集标记信息的时间点

        self.video = video
        self.rgb_arr = self.video.get_frame(self.point_time)
        self.boundary_marker = None         # 检测范围
        self.back_color = {}                # 背景颜色信息

        self.get_boundary_marker()
        self.get_back_color()

    def get_boundary_marker(self):
        # 获取画面边界
        print("开始获取画面边界...")
        top_range = (0, (self.video.h // 3))                    # 画面上半部分的行数范围
        bottom_range = ((self.video.h // 3 * 2), self.video.h)  # 画面下半部分的行数范围
        left_range = (0, (self.video.w // 3))                   # 画面左半部分的列数范围
        right_range = ((self.video.w // 3 * 2), self.video.w)   # 画面右半部分的列数范围

        self.boundary_marker = {
            "top": self.get_most_black(top_range, "line"),
            "bottom" : self.get_most_black(bottom_range, "line"),
            "left": self.get_most_black(left_range, "row"),
            "right": self.get_most_black(right_range, "row")
        }
        print("画面边界为", self.boundary_marker)

    def get_most_black(self, my_range, key):
        # 返回range范围内黑点数最多的行或列id(key用来标记是行还是列)
        start_id, end_id = my_range
        max_black_dot_sum = 0
        boundary_marker_row = 0
        
        for id in range(start_id, end_id):
            line_or_row = self.rgb_arr[id] if key == "line" else self.traversal_row(id)
            black_dot_sum = self.get_black_dot_sum(line_or_row)
            if black_dot_sum > max_black_dot_sum:
                max_black_dot_sum = black_dot_sum
                boundary_marker_row = id
        return boundary_marker_row

    def traversal_row(self, row_num):
        # 用来按列遍历rgb数组的生成器
        for i in range(self.video.h):
            rgb_dot = self.rgb_arr[i][row_num]
            yield rgb_dot      

    def get_black_dot_sum(self, one_line_rgb_li):
        # 返回某一行或列里黑点的个数
        black_dot_sum = 0
        for rgb in one_line_rgb_li:
            if if_black(rgb):
                black_dot_sum += 1
        return black_dot_sum

    def get_back_color(self):
        print("获取背景颜色信息...")
        rgb_li = self.rgb_arr[self.boundary_marker["top"]]
        hub_li = []
        saturation_li = []
        for rgb in rgb_li:
            if not if_black(rgb):
                h, s, _ = rgb2hsb(rgb)
                hub_li.append(h)
                saturation_li.append(s)
        self.back_color["h"] = get_most_value(hub_li)
        self.back_color["s"] = get_most_value(saturation_li)
        print("背景颜色信息为", self.back_color)