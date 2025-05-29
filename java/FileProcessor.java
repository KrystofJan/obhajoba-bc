import java.io.*;
import java.util.regex.*;

public class FileProcessor {
    public static void filterLinesByRegex(String inputFilePath, String regexPattern, String outputFilePath) {
        BufferedReader reader = null;
        BufferedWriter writer = null;
        int lineNumber = 0;

        try {
            reader = new BufferedReader(new FileReader(inputFilePath));
            writer = new BufferedWriter(new FileWriter(outputFilePath));
            Pattern pattern = Pattern.compile(regexPattern);

            String line;
            while ((line = reader.readLine()) != null) {
                lineNumber++;
                try {
                    Matcher matcher = pattern.matcher(line);
                    if (matcher.find()) {
                        writer.write(line);
                        writer.newLine();
                    }
                } catch (Exception e) {
                    System.err.printf("⚠️ Chyba při zpracování regulárního výrazu na řádku %d: %s%n", lineNumber, e.getMessage());
                }
            }

        } catch (FileNotFoundException e) {
            System.err.printf("❌ Soubor '%s' nebyl nalezen.%n", inputFilePath);
        } catch (IOException e) {
            System.err.printf("❌ IO chyba: %s%n", e.getMessage());
        } catch (OutOfMemoryError e) {
            System.err.println("❌ Nedostatek paměti při zpracování souboru.");
        } catch (Exception e) {
            System.err.printf("❌ Neočekávaná chyba: %s%n", e.getMessage());
        } finally {
            if (reader != null) {
                try { reader.close(); } catch (IOException e) {
                    System.err.printf("⚠️ Chyba při zavírání vstupního souboru: %s%n", e.getMessage());
                }
            }
            if (writer != null) {
                try { writer.close(); } catch (IOException e) {
                    System.err.printf("⚠️ Chyba při zavírání výstupního souboru: %s%n", e.getMessage());
                }
            }
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
