class HashMap:
  def __init__(self, s=0):
    self.a = {}

  def add(self, key, value):
    self.a[key] = value

  def get(self, key):
    if key not in self.a:
      return -1
    return self.a[key]

def finder(nums, target):
  x = HashMap()

  for i in range(0, len(nums)):
    required = target - nums[i]
    pos = x.get(required)
    if(pos == -1):
      x.add(nums[i], i)
      continue
    
    return [pos, i]

inputs = [0,3,-3,4,-1]
required = -1

needed_result = -1
print(finder(inputs, needed_result ))

