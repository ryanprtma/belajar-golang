function repeatString(s,n){
    let arr = []
    let count = 0
    let i = 0
    
    if(i==0){
        for(i;i<=s.length-1;i++){
            arr.push(s[i])
            count++
        }
    }
    

    if(count==s.length){
        count=0
        i=0
    }

    console.log(i)
    return arr
}

console.log(repeatString("abcabcabc", 20))