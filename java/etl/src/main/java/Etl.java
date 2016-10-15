import java.util.List;
import java.util.Map;
import java.util.HashMap;

public class Etl {
    public Map<String, Integer> transform(Map<Integer, List<String>> old) {
        Map<String, Integer> improved = new HashMap<String, Integer>();
        old.forEach((score, letters) ->
            letters.forEach(letter ->
                improved.put(letter.toLowerCase(), score)
            )
        );
        return improved;
    }
}
