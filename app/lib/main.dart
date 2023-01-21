import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}
final List<String> entries = <String>['Recipe A', 'Recipe B', 'Recipe C','Recipe D'];
//final List<int> colorCodes = <int>[600, 500, 400,300];
final List<String> images = <String>['images/shrimp.jpg','images/sushi.jpg','images/lasagna.jpg','images/empanadas.jpg'];

//class containing basic information of each post
class Post{
  String title;
  String image;
  int pID;
  Post(this.title,this.image, this.pID);
}

//final List<Post> postList = <Post>[];

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Palette Pal',
      theme: ThemeData(
        // This is the theme of your application.
        //
        // Try running your application with "flutter run". You'll see the
        // application has a blue toolbar. Then, without quitting the app, try
        // changing the primarySwatch below to Colors.green and then invoke
        // "hot reload" (press "r" in the console where you ran "flutter run",
        // or simply save your changes to "hot reload" in a Flutter IDE).
        // Notice that the counter didn't reset back to zero; the application
        // is not restarted.
        primarySwatch: Colors.blue,
      ),
      home: const MyHomePage(title: 'Palette Pal'),
    );

  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key, required this.title}) : super(key: key);

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _counter = 0;

  void _incrementCounter() {
    setState(() {
      // This call to setState tells the Flutter framework that something has
      // changed in this State, which causes it to rerun the build method below
      // so that the display can reflect the updated values. If we changed
      // _counter without calling setState(), then the build method would not be
      // called again, and so nothing would appear to happen.
      _counter++;
    });
  }


  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done
    // by the _incrementCounter method above.
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.
    return Scaffold(
      appBar: AppBar(
        // Here we take the value from the MyHomePage object that was created by
        // the App.build method, and use it to set our appbar title.
        title: Text(widget.title),
      ),
      body: ListView.separated(

        padding: const EdgeInsets.all(8),
        itemCount: entries.length,
        itemBuilder:(BuildContext context, int index) {
          return Container(

            height: 350,
            width: 80,
            //color: Colors.amber[colorCodes[index]],
            decoration: BoxDecoration(
              image: DecorationImage(
                image: AssetImage(
                  '${images[index]}'),
                fit: BoxFit.fill,
              ),
              shape: BoxShape.rectangle,
            ),
            //child: Center(
                child: Text('${entries[index]}', textAlign: TextAlign.start)
               // child: Image(image: AssetImage('${images[index]}'))
            //),
          );
        },
        separatorBuilder: (BuildContext context, int index) => const Divider(),
      ),
        // Center is a layout widget. It takes a single child and positions it
        // in the middle of the parent.

      bottomNavigationBar: BottomNavigationBar(
        type: BottomNavigationBarType.fixed,
        //backgroundColor: Colors.blueAccent,
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.search),
            label: 'Search',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.add_circle_outline),
            label: 'Post',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.account_circle),
            label: 'Account',
          ),
        ],
      ),
      /*
      floatingActionButton: FloatingActionButton(
        onPressed: _incrementCounter,
        tooltip: 'Increment',
        child: const Icon(Icons.safety_check),
      ),*/ // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
