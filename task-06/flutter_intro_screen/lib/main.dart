import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:introduction_screen/';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: IntroductionPage(),
    );
  }
}

class IntroductionPage extends StatelessWidget {
  void WelcomePage(context) {
    Navigator.pop(context);
    Navigator.of(context).push(MaterialPageRoute(builder: (_) => HomePage()))
  }
}

class HomePage extends StatelessWidget{
  Widget build(BuildContext context){
    return Scaffold(
      home:pages(),

    )
  }
}

List<PageViewModel> Pages() {
    return [
      PageViewModel(
          title: "YOGA Surge",
          body: "Welcome to Yoga World,Inhale the Future,Exhale the past.",
          image:
              Center(child: Image.asset("./assets/page1.png", height: 175.0))),
      PageViewModel(
          title: "pizza Freaks Exercises",
          body: "Letting go is the hardest asana",
          image: Center(child: Image.asset("assets/page2.png", height: 175.0))),
      PageViewModel(
          title: "Cycling",
          body: "You cannot always control what "
              "goes on outside. But you can always "
              "control what goes on inside",
          image: Center(child: Image.asset("assets/page3.png", height: 175.0))),
      PageViewModel(
          title: "MEDITATION",
          body: "The longest journey of any person is the journey inward",
          image: Center(child: Image.asset("assets/page4.png", height: 175.0)))
    ];
  }