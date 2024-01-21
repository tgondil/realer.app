import "./messageBar.css";
import Grid from "@mui/material/Grid";
import DefaultScreen from "./default";

const MessageBar = () => {
  return (
    <>
      <DefaultScreen></DefaultScreen>
      <Grid
        style={{
          paddingRight: "0",
          margin: "10px",
          background: "white",
          borderRadius: "50px 50px 0px 0px",
          width: "100%",
        }}
        className="bar"
        container
        spacing={2}
      ></Grid>
    </>
  );
};

export default MessageBar;
