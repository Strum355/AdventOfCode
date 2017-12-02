import com.opencsv.CSVParserBuilder
import com.opencsv.CSVReaderBuilder
import java.io.*
import java.util.*

fun main(args: Array<String>) {
    val parser = CSVParserBuilder().withSeparator('\t').withIgnoreQuotations(true).build()
    val reader = CSVReaderBuilder(FileReader("../day2_input.txt")).withCSVParser(parser).build()
    var total = 0
    while(true) {
        val row = reader.readNext()?.map { it.toInt() }?.toTypedArray() ?: break
        row.forEachIndexed { i, num ->
            val slice = row.sliceArray(IntRange(i+1, row.size-1))
            val res = findDivisible(num, slice)
            total += if(res != null) res else 0
        }
    }
    println(total)
}

fun findDivisible(i: Int, s: Array<Int>): Int? {
    s.forEach{
        if(i%it == 0){
            return i/it
        }
        if(it%i == 0){
            return it/i
        }
    }
    return null
}