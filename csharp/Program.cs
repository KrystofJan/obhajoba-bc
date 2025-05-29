using System;
using System.IO;
using System.Text.RegularExpressions;

class FileProcessor {
    public static void FilterLinesByRegex(string inputFilePath, string regexPattern, string outputFilePath) {
	StreamReader reader = null;
	StreamWriter writer = null;
	int lineNumber = 0;

	try {
	    reader = new StreamReader(inputFilePath);
	    writer = new StreamWriter(outputFilePath);
	    Regex regex = new Regex(regexPattern);

	    string line;
	    while ((line = reader.ReadLine()) != null) {
		lineNumber++;
		try {
		    if (regex.IsMatch(line)) {
			writer.WriteLine(line);
		    }
		}
		catch (Exception ex) {
		    Console.WriteLine($"⚠️ Chyba při zpracování regulárního výrazu na řádku {lineNumber}: {ex.Message}");
		}
	    }
	}
	catch (FileNotFoundException) {
	    Console.WriteLine($"❌ Soubor '{inputFilePath}' nebyl nalezen.");
	}
	catch (IOException ex) {
	    Console.WriteLine($"❌ IO chyba: {ex.Message}");
	}
	catch (OutOfMemoryException) {
	    Console.WriteLine("❌ Nedostatek paměti při zpracování souboru.");
	}
	catch (Exception ex) {
	    Console.WriteLine($"❌ Neočekávaná chyba: {ex.Message}");
	}
	finally {
	    if (reader != null) {
		try { 
		    reader.Close(); 
		} catch (Exception ex) {
		    Console.WriteLine($"⚠️ Chyba při zavírání vstupního souboru: {ex.Message}"); 
		}
	    }

	    if (writer != null) {
		try { 
		    writer.Close(); 
		} catch (Exception ex) {
		    Console.WriteLine($"⚠️ Chyba při zavírání výstupního souboru: {ex.Message}"); 
		}
	    }
	}
    }
}

class Program {
    static void Main(string[] args) {
        if (args.Length != 3) {
            Console.WriteLine("❗ Použití: <vstupní_soubor> <výstupní_soubor> <regulární_výraz>");
            return;
        }

        string inputFile = args[0];
        string outputFile = args[1];
        string pattern = args[2];

        Console.WriteLine($"📂 Zpracovávám soubor '{inputFile}' s výrazem '{pattern}'");
        FileProcessor.FilterLinesByRegex(inputFile, pattern, outputFile);
        Console.WriteLine("✅ Hotovo. Výsledky jsou zapsány do souboru.");
    }
}
