import Box from "@mui/material/Box";
import { Typography } from "@mui/material";
import { useMemo } from "react";
import { DateTime } from "luxon";

interface Props {
  title: string;
  date: string;
  hyperlink: string;
}

export const ArticleCard = (props: Props) => {
  const { title, date, hyperlink } = props;

  const formattedDate = useMemo(() => {
    const truncatedIso = date.split(".")[0] + "Z";
    return DateTime.fromISO(truncatedIso).toFormat("LLLL d, yyyy");
  }, [date]);

  const handleClick = () => {
    window.open(hyperlink, "_blank", "noopener,noreferrer");
  };
  return (
    <Box
      padding={4}
      sx={{
        fontWeight: "bold",
        color: "white",
        textTransform: "none",
        borderRadius: 0,
        boxShadow: "3px 3px 0px 0px black",
        transition: "transform 0.2s ease, box-shadow 0.2s ease",
        border: "1px solid black",
        position: "relative",
        zIndex: 2,
        "&:hover": {
          transform: "translateY(-4px)",
          boxShadow: "8px 8px 0px 0px black",
          cursor: "pointer",
        },
      }}
    >
      <Typography
        variant={"h5"}
        fontWeight={"bold"}
        component={"a"}
        onClick={handleClick}
        sx={{
          "&:hover": {
            cursor: "pointer",
          },
        }}
      >
        {title}
      </Typography>
      <Typography variant={"subtitle1"}>{formattedDate}</Typography>
    </Box>
  );
};
