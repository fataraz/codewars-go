package kyuu8

//The function is not returning the correct values. Can you figure out why?
//
//Example (Input --> Output ):
//
//3 --> "Earth"

func GetPlanetName(ID int) string {
	arr := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto"}
	return arr[ID-1]
}
