import { Component, OnInit } from "@angular/core";
import { ActivatedRoute } from "@angular/router";

@Component({
  selector: "app-card",
  templateUrl: "./card.component.html",
  styleUrls: ["./card.component.scss"]
})
export class CardComponent implements OnInit {
  imagePath = "";

  constructor(private route: ActivatedRoute) {}

  ngOnInit() {
    let fieldId = this.route.snapshot.paramMap.get("path");
    this.imagePath = `download/${fieldId}.png`;
    console.log("ngOnInit", this.imagePath);
  }
}
