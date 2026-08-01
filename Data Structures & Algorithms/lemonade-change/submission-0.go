func lemonadeChange(bills []int) bool {
    five, ten := 0, 0
    // result := false
	for _, v := range bills{
        if v == 5{
            five ++
        }else if(v ==10){
            if five >0{
                five --
                ten++
            }else{
                return false
            }
        }else{ // v=20
            if ten >0 && five >0 {
                ten --
                five --
            }else if five >2{
                five --
                five --
                five --
            }else{
                return false
            }
        }
    }
    return true
}
// 5 20 10 5
// 5 20 -> not have enough change ->false

//5 10 5 5 20
//5 
//10 5