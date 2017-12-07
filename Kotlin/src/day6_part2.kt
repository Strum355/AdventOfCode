import java.io.File
import java.util.*

fun main(a: Array<String>) {
    val inputStream = File("../day6_input.txt").inputStream()
    val input = inputStream.bufferedReader().use { it.readText() }

    var blocks = input.split('\t').map { it.toInt() } as ArrayList<Int>

    var jumps = 1
    val state = linkedMapOf<String, Boolean>()

    while(true){
        blocks = distribute(blocks)
        if(state.containsKey(blocks.toString())){
            var i = 1
            state.forEach loop@ { s, _ ->
                if(s == blocks.toString()){
                    println(jumps-i)
                    return@loop
                }
                i++
            }
            return
        }
        state.put(blocks.toString(), true)
        jumps++
    }
}