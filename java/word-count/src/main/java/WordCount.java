import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class WordCount {

    public Map<String, Integer> phrase(String given) {
        List<String> elements = Arrays.asList(given.toLowerCase().split("\\W+"));
        Map<String, Integer> wordCount = elements
            .stream()
            .collect(Collectors.groupingBy(s -> s, Collectors.summingInt(s -> 1)));
        return wordCount;
    }
}
