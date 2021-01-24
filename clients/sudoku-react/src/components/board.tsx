import * as React from "react";
import "../App.css";

export interface BoardProps {
  content: number[][];
}

export interface BoardState {
  content: number[][];
}

class Board extends React.Component<BoardProps, BoardState> {
  state = { content: [] };

  render() {
    const { content } = this.props;
    return <div className="board">{this.getCoordinateFields(content)}</div>;
  }

  getCoordinateFields = (content: number[][]) => {
    return content.map((row, x) => {
      return (
        <div>
          {row.map((col, y) => {
            return <div className="coordinate">{this.formatCell(col)}</div>;
          })}
        </div>
      );
    });
  };

  formatCell = (val: number) => {
    if (val > 0) {
      return <span>{val}</span>;
    }

    return <input type="text" className="coordinate-editable " />;
  };
}

export default Board;
