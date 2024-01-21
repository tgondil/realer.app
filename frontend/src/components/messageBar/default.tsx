import "./default.css";

const DefaultScreen = () => {
  return (
    <>
      <div
        style={{
          height: "90vh",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          flexWrap: "wrap",
          rowGap: "0",
        }}
      >
        <h1 className="hero big gradient"> realer.app </h1>
        <br />
      </div>
    </>
  );
};

export default DefaultScreen;
