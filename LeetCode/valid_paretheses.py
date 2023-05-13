class Solution:
    def isValid(self, s: str) -> bool:
        str_valid = True
        brackets_count = {
            '{}': 0, 
            '()': 0, 
            '[]': 0
        }
        
        for char in s:
            if char == '{':
                brackets_count['{}'] += 1
                
            elif char == '}':
                brackets_count['{}'] -= 1
                if brackets_count['{}'] < 0:
                    str_valid = False
                    return str_valid
                
            elif char == '(':
                brackets_count['()'] += 1
                
            elif char == ')':
                brackets_count['()'] -= 1
                if brackets_count['()'] < 0:
                    str_valid = False
                    return str_valid
                
            elif char == '[':
                brackets_count['[]'] += 1
                
            elif char == ']':
                brackets_count['[]'] -= 1
                if brackets_count['[]'] < 0:
                    str_valid = False
                    return str_valid
                
        return str_valid
                
            