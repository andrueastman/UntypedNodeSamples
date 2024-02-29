/*
 * This Java source file was generated by the Gradle 'init' task.
 */
package javaclient.app;

import javaclient.list.LinkedList;

import static javaclient.utilities.StringUtils.join;
import static javaclient.utilities.StringUtils.split;
import static javaclient.app.MessageUtils.getMessage;

import org.apache.commons.text.WordUtils;

public class App {
    public static void main(String[] args) {
        LinkedList tokens;
        tokens = split(getMessage());
        String result = join(tokens);
        System.out.println(WordUtils.capitalize(result));
    }
}
