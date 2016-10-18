import java.util.HashMap;
import java.util.Map;

public class DNA {
    private Map<Character, Integer> nucleotides;

    public Integer count(char nucleotide) throws IllegalArgumentException {
        if (!nucleotides.containsKey(nucleotide)) throw new IllegalArgumentException();
        return nucleotides.get(nucleotide);
    }

    public Map<Character, Integer> nucleotideCounts() {
        return nucleotides;
    }

    public DNA(String DNAString) {
        nucleotides = new HashMap<Character, Integer>();
        nucleotides.put('A', 0);
        nucleotides.put('C', 0);
        nucleotides.put('G', 0);
        nucleotides.put('T', 0);
        for (int i = 0; i < DNAString.length(); i++) {
            nucleotides.merge(DNAString.charAt(i), 1, Integer::sum);
        }
    }
}
