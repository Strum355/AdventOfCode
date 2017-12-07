import java.io.File

fun main(args: Array<String>) {
    var total = 0
    val inputStream = File("../day1_input.txt").inputStream()
    val s = inputStream.bufferedReader().use { it.readText() }

    s.forEachIndexed { i, j ->
        if (j == s.get((i+1)%s.length)){
            total += j.toInt()-48
        }
    }
    println(total)
}