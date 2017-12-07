import java.io.File
import java.io.InputStream

fun main(args: Array<String>) {
    var pos = 0
    var jumps = 0

    val inputStream = File("../day5_input.txt").inputStream()
    val input = inputStream.bufferedReader().use { it.readText() }
    
    var instructions = input.split('\n').map{ it.toInt() } as ArrayList<Int>

    while(pos < instructions.size){
        val jumpAmount = instructions.get(pos)
        instructions.set(pos, jumpAmount+fun(): Int { if(jumpAmount > 2) return -1 else return 1 }.invoke())
        pos += jumpAmount
        jumps++
    }
    println(jumps)
}