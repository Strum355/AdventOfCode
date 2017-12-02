import com.opencsv.CSVParserBuilder
import com.opencsv.CSVReaderBuilder
import java.io.*
import java.util.*

fun main(args: Array<String>) {
    val parser = CSVParserBuilder().withSeparator('\t').withIgnoreQuotations(true).build()
    val reader = CSVReaderBuilder(FileReader("../day2_input.txt")).withCSVParser(parser).build()
    var total = 0
    while(true) {
        val row = reader.readNext()?.map { it.toInt() } ?: break
        val max = row.max()
        val min = row.min()
        total += max!!-min!!
    }
    println(total)
}