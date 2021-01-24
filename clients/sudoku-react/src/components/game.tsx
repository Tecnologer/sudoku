import * as React from "react";
import "../App.css";
import Board from "./board";
export interface GameProps {}

export interface GameState {}

class Game extends React.Component<GameProps, GameState> {
  state = {};
  render() {
    let content = this.getContent();
    return (
      <div className="game">
        <Board content={content} />
      </div>
    );
  }

  getContent = () => {
    let board: string[][] = [
      ["3", "", "", "5", "", "", "", "6", "9"],
      ["4", "", "2", "", "", "", "", "", ""],
      ["", "", "5", "", "6", "", "8", "7", ""],
      ["", "", "", "1", "", "2", "", "", "7"],
      ["", "", "1", "", "", "", "3", "", ""],
      ["7", "", "", "9", "", "4", "", "", ""],
      ["", "2", "9", "", "1", "", "6", "", ""],
      ["", "", "", "", "", "", "4", "", "8"],
      ["5", "3", "", "", "", "6", "", "", "2"],
    ];
    return board;
  };
}

export default Game;
