import Box from "@mui/material/Box";
import {Typography} from "@mui/material";

export const ArticleCard = () => {
    return (<Box padding={4} sx={{
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
            // background: "linear-gradient(to right, #6366f1, #7c3aed)",
            cursor: "pointer",
        },
    }}>
        <Typography variant={'h5'} fontWeight={'bold'}>Title</Typography>
        <Typography variant={'subtitle1'}>Date</Typography>
    </Box>)
}