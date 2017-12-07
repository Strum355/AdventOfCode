import java.io.File

fun main(a: Array<String>) {
    val inputStream = File("../day6_input.txt").inputStream()
    val input = inputStream.bufferedReader().use { it.readText() }

    var blocks = input.split('\t').map { it.toInt() } as ArrayList<Int>

    var jumps = 1
    val state = mutableMapOf<String, Boolean>()

    while(true){
        blocks = distribute(blocks)
        println(blocks)
        if(state.containsKey(blocks.toString())){
            println(jumps)
            return
        }
        state.put(blocks.toString(), true)
        jumps++
    }
}

fun distribute(a: ArrayList<Int>): ArrayList<Int> {
    var (max, index) = max(a)
    a[index] = 0
    while(max != 0) {
        index = (index + 1)%a.size
        a[index]++
        max--
    }
    return a
}

fun max(a: ArrayList<Int>): Pair<Int, Int> {
    var max = a.get(0)
    var index = 0
    a.forEachIndexed { i, v ->
        if(v > max){
            max = v
            index = i
        }
    }
    return Pair(max, index)
}