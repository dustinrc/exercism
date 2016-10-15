import java.util.List;
import java.util.Map;
import java.util.HashMap;

public class Etl {
    public Map<String, Integer> transform(Map<Integer, List<String>> old) {
        Map<String, Integer> improved = new HashMap<String, Integer>();
        for (Integer score : old.keySet()) {
            for (String letter : old.get(score)) {
                improved.put(letter.toLowerCase(), score);
            }
        }
        return improved;
    }
}
