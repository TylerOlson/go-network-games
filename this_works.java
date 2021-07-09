public class this_works {

    TerminalScreen currentScreen;

    public this_works() {
        TerminalScreen mainMenuScreen = new MainMenuScreen("Main Menu", 50, 50);
        TerminalScreen gameScreen = new GameScreen("Game", 50, 50);

        currentScreen = mainMenuScreen;
        draw();
        resizeEvent();
        draw();
        currentScreen = gameScreen;
        draw();
    }

    public void draw() {
        currentScreen.DrawContent();
        System.out.println("");
    }

    public void resizeEvent() { // this occurs during resize
        currentScreen.UpdateSize(20, 30);
    }

    public static void main(String[] args) throws Exception {
        new this_works();
    }
}

class TerminalScreen {

    private String name;
    private int width, height;

    public TerminalScreen(String name, int width, int height) {
        this.name = name;
        this.width = width;
        this.height = height;
    }

    public void DrawContent() {
        System.out.println("     " + name + " " + width + "," + height);
    }

    public void KeyEvent(char key) {
    }

    public void UpdateSize(int w, int h) {
        this.width = w;
        this.height = h;
    }
}

class MainMenuScreen extends TerminalScreen {

    private int currentSelection = 0;
    private int menuItems = 2;

    public MainMenuScreen(String name, int width, int height) {
        super(name, width, height);
    }

    public void DrawContent() {
        super.DrawContent();
        if (currentSelection == 0) {
            System.out.println("1. Start Game <-");
        } else {
            System.out.println("1. Start Game");
        }

        if (currentSelection == 1) {
            System.out.println("2. Exit <-");
        } else {
            System.out.println("2. Exit");
        }

        // obv this logic is really bad aand I would make it better blah lbah
    }

    public void KeyEvent(char key) {
        // if w
        if (key == 'w') {
            currentSelection--;
        } else if (key == 's') {
            currentSelection++;
        }

        if (currentSelection < 0) {
            currentSelection = menuItems - 1;
        } else if (currentSelection >= menuItems) {
            currentSelection = 0;
        }
    }

}

class GameScreen extends TerminalScreen {

    public GameScreen(String name, int width, int height) {
        super(name, width, height);
    }

    public void DrawContent() {
        super.DrawContent();
        System.out.println("HERE IS MY GAME");
    }

    public void KeyEvent(char key) {
    }
}
