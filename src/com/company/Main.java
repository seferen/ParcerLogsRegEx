package com.company;

import java.io.BufferedWriter;
import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.nio.file.*;
import java.nio.file.attribute.BasicFileAttributes;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 *2. Написать парсер логов, который выполняет следующие действия над логами в п.1:
 *    а) С помощью регулярного выражения задать фразу поиска и все найденные строки записать в новый файл
 *    а) 3 параметра: <фраза_поиска>,<путь_до_лог_файла(ов)>,<имя_нового_файла>
 */

public class Main {

    public static void main(String[] args) {

        String rregEx = args[0]; //принимаемое регулярное выражение
        String path = args[1]; //принимаемый путь до файлов или папок
        String fileResult = args[2]; //имя нового файла

        List<String> result = new ArrayList<>();

        String string = "mama1234151, papa15345134, and I was born in the Kineshma city";

        try(BufferedWriter bf = new BufferedWriter(new FileWriter(new File(fileResult)))) {

            Files.walkFileTree(Paths.get(path), new SimpleFileVisitor<>() {
                @Override
                public FileVisitResult visitFile(Path file, BasicFileAttributes attrs) throws IOException {
                    for (String s : Files.readAllLines(file)) {
                        Matcher matcher = Pattern.compile(rregEx).matcher(s);
                        while (matcher.find()) {

                            for (int i = 0; i <= matcher.groupCount(); i++) {

                                bf.write(matcher.group(i) + ",");

                            }
                            bf.newLine();

                        }
                    }

                    return FileVisitResult.CONTINUE;
                }
            });
        }catch (IOException e){
            System.out.print("Произошлк какая то фигня возможно файл не полный");
        }


	// write your code here
    }
}
