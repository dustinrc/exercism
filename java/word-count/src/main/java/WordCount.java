import java.util.Map;
import java.util.HashMap;

public class WordCount {
    private Map<String, Integer> wordCount;

    public Map<String, Integer> phrase(String given) {
        for ( String word : given.split("\\W+") ) {
            wordCount.merge(word.toLowerCase(), 1, Integer::sum);
        }
        return wordCount;
    }

    public WordCount() {
        wordCount = new HashMap<String, Integer>();
    }
}
