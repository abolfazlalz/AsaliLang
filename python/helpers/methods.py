import math


class Methods:
    def print(self, var: list):
        result: list[str] = []
        for item in var:
            result.append(str(item))

        print(" ".join(result))

    def sin(self, val: list):
        if len(val) < 1:
            assert "required one parameter"
            return
        return math.sin(val[0])
