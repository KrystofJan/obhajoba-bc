use std::env;
use std::fs::File;
use std::io::{self, BufRead, BufReader, BufWriter, Write};
use regex::Regex;

#[allow(clippy::cognitive_complexity)]
fn filter_lines_by_regex(input_file: &str, pattern: &str, output_file: &str) -> io::Result<()> {
    let reader = BufReader::new(File::open(input_file)?);
    let mut writer = BufWriter::new(File::create(output_file)?);
    let re = Regex::new(pattern).map_err(|e| {
        eprintln!("❌ Chyba v regulárním výrazu: {}", e);
        io::Error::new(io::ErrorKind::InvalidInput, "Neplatný regex")
    })?;

    for (line_number, line_result) in reader.lines().enumerate() {
        let line = match line_result {
            Ok(content) => content,
            Err(e) => {
                eprintln!("❌ Chyba při čtení řádku {}: {}", line_number + 1, e);
                continue;
            }
        };

        if re.is_match(&line) {
                if let Err(e) = writeln!(writer, "{}", line) {
                    eprintln!("❌ Chyba při zápisu na řádku {}: {}", line_number + 1, e);
                    return Err(e);
                }
            }
    }

    writer.flush()?;
    Ok(())
}

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() != 4 {
        eprintln!("❗ Použití: <vstupní_soubor> <výstupní_soubor> <regulární_výraz>");
        return;
    }

    let input_file = &args[1];
    let output_file = &args[2];
    let pattern = &args[3];

    println!("📂 Zpracovávám soubor '{}' s výrazem '{}'", input_file, pattern);

    if let Err(e) = filter_lines_by_regex(input_file, pattern, output_file) {
        eprintln!("❌ Chyba během zpracování: {}", e);
    } else {
        println!("✅ Hotovo. Výsledky jsou zapsány do souboru.");
    }
}
