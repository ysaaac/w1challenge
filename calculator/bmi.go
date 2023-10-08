package calculator

func CalculateBMI(weight, height float32) string {
	result := weight / (height * height)

	if result < 18.5 {
		return "You are underweight, add more potato to the broth"
	} else if result < 25 {
		return "You have a normal weight, I have healthy envy of you"
	} else {
		return "You are overweight, I know, the pandemic has affected us all"
	}
}

