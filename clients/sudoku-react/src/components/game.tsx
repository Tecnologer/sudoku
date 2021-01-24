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
    let board: number[][] = [
      [3, 0, 0, 5, 0, 0, 0, 6, 9],
      [4, 0, 2, 0, 0, 0, 0, 0, 0],
      [0, 0, 5, 0, 6, 0, 8, 7, 0],
      [0, 0, 0, 1, 0, 2, 0, 0, 7],
      [0, 0, 1, 0, 0, 0, 3, 0, 0],
      [7, 0, 0, 9, 0, 4, 0, 0, 0],
      [0, 2, 9, 0, 1, 0, 6, 0, 0],
      [0, 0, 0, 0, 0, 0, 4, 0, 8],
      [5, 3, 0, 0, 0, 6, 0, 0, 2],
    ];
    return board;
  };
}

export default Game;
