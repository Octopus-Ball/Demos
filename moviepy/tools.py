def rgb2hsb(rgb):
        # 将RGB值转换为HSB
        r, g, b = rgb[0]/255.0, rgb[1]/255.0, rgb[2]/255.0
        mx = max(r, g, b)
        mn = min(r, g, b)
        m = mx-mn
        if mx == mn:
            h = 0
        elif mx == r:
            if g >= b:
                h = ((g-b)/m)*60
            else:
                h = ((g-b)/m)*60 + 360
        elif mx == g:
            h = ((b-r)/m)*60 + 120
        elif mx == b:
            h = ((r-g)/m)*60 + 240
        if mx == 0:
            s = 0
        else:
            s = m/mx
        v = mx
        s *= 100
        v *= 100
        # return h, s, v
        return int(h), int(s), int(v)

def if_black(rgb):
        # 传入RGB，判断是否是黑色
        threshold_value = 70
        for v in rgb:
            if v > threshold_value:
                return False
        return True

def get_most_value(li):
    # 根据传入的列表找出其中的众数
        most_sum = 0
        most_v = None
        v_dic = {}
        # 统计各hub_v的个数
        for v in li:
            if v in v_dic:
                v_dic[v] += 1
            else:
                v_dic[v] = 1
        # 找出数量最大的hub_v
        for v, sum in v_dic.items():
            if sum > most_sum:
                most_sum = sum
                most_v = v
        
        return most_v