def isPalindrome(self, x):
  strX = str(x)
  if strX[::-1] == strX:
      return True
  return False