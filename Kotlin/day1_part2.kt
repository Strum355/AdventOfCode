import java.io.File
import java.io.InputStream

fun main(args: Array<String>) {
    var total = 0
    val inputStream = File("../day1_input.txt").inputStream()
    val s = inputStream.bufferedReader().use { it.readText() }
    s.forEachIndexed { i, j ->
        if (j == s.get((i+s.length/2)%s.length)){
            total += j.toInt()-48
        }
    }
    println(total)
}