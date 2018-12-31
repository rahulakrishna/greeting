import { Component, OnInit } from "@angular/core";
import { Router } from "@angular/router";

import { Poem } from "../poem";
import { PoemService } from "../poem.service";

@Component({
  selector: "app-home",
  templateUrl: "./home.component.html",
  styleUrls: ["./home.component.scss"]
})
export class HomeComponent implements OnInit {
  poem = new Poem();
  body = "";
  constructor(private poemService: PoemService, private router: Router) {}

  ngOnInit(): void {
    this.poem.body.map(line => (this.body += line + "\n"));
  }

  generate(): void {
    if (Object.keys(this.poem).length !== 0) {
      this.poem.body = this.body.split("\n");
      this.poemService.generate(this.poem).then(resp => {
        console.log(resp);
        this.router.navigate(["/card", resp.fileId]);
      });
    } else {
      console.log("Empty!");
    }
  }
}
