fun main(args: Array<String>) {
    var total = 0
    val s = "1234" //input number
    s.forEachIndexed { i, j ->
        if (j == s.get((i+s.length/2)%s.length)){
            total += j.toInt()-48
        }
    }
    println(total)
}