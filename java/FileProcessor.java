import java.io.*;
import java.util.regex.*;

public class FileProcessor {
    public static void filterLinesByRegex(String inputFilePath, String regexPattern, String outputFilePath) {
        Pattern pattern = Pattern.compile(regexPattern);
        try (BufferedReader reader = new BufferedReader(new FileReader(inputFilePath));
             BufferedWriter writer = new BufferedWriter(new FileWriter(outputFilePath))) {
            String line;
            int lineNumber = 0;
            while ((line = reader.readLine()) != null) {
                lineNumber++;
                try {
                    if (pattern.matcher(line).find()) { // +1
                        writer.write(line);
                        writer.newLine();
                    }
                } catch (Exception e) { // +1
                    System.err.printf("⚠️ Chyba při zpracování regulárního výrazu na řádku %d: %s%n", lineNumber, e.getMessage());
                }
            }
        } catch (FileNotFoundException e) { // +1
            System.err.printf("❌ Soubor '%s' nebyl nalezen.%n", inputFilePath);
        } catch (IOException e) { // +1
            System.err.printf("❌ IO chyba: %s%n", e.getMessage());
        } catch (OutOfMemoryError e) { // +1
            System.err.println("❌ Nedostatek paměti při zpracování souboru.");
        } catch (Exception e) { // +1
            System.err.printf("❌ Neočekávaná chyba: %s%n", e.getMessage());
        }
    }
    public static void main(String[] args) {
        if (args.length != 3) {
            System.err.println("❗ Použití: <vstupní_soubor> <výstupní_soubor> <regulární_výraz>");
            return;
        }

        String inputFile = args[0];
        String outputFile = args[1];
        String pattern = args[2];

        System.out.printf("📂 Zpracovávám soubor '%s' s výrazem '%s'%n", inputFile, pattern);
        filterLinesByRegex(inputFile, pattern, outputFile);
        System.out.println("✅ Hotovo. Výsledky jsou zapsány do výstupního souboru.");
    }
}
