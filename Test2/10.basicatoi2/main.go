package main

func BasicAtoi2(s string) int {
result :=0
for i:=0;i < len(s);i++{
	if s[i] >= '0' && s[i] <='9' {
		result = result*10 + int(s[i]-'0')
	
	}else{
		return 0
	}
}
return result
}
