import { Typography } from "@mui/material"
import {ArticleCard} from "../../components/ArticleCard.tsx";

export const NewsFeed = () => {
    return <>
        <Typography variant={'h4'} textAlign={'center'} sx={{textDecoration: "underline"}} py={4}>Latest Articles
        </Typography>
        <ArticleCard />
    </>
}