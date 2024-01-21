import { lime } from "@mui/material/colors";

import GroupsIcon from "@mui/icons-material/Groups";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import { createTheme } from "@mui/material/styles";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";

import "./topBar.css";

const theme = createTheme({
  palette: {
    primary: lime,
    secondary: {
      main: "#E0C2FF",
      light: "#F5EBFF",
      // dark: will be calculated from palette.secondary.main,
      contrastText: "#47008F",
    },
  },
});

function topBar() {
  return (
    <div className="topBar">
      <GroupsIcon
        className="gradient"
        style={{ fontSize: 40, paddingTop: "7px", color: "white" }}
      />{" "}
      <MoreVertIcon
        color="primary"
        style={{ fontSize: 40, paddingTop: "7px", color: "white" }}
      />
      <AccountCircleIcon
        color="primary"
        style={{ fontSize: 40, paddingTop: "7px", color: "white" }}
      />
    </div>
  );
}

export default topBar;
