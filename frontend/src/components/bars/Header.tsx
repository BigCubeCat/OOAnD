import {Box, IconButton, Typography} from "@mui/joy";
import PersonOutlineIcon from '@mui/icons-material/PersonOutline';

export default function Header() {
    return (
        <Box sx={{
            bgcolor: 'primary.700',
            '&:hover': {
                bgcolor: 'primary.400',
            },
            padding: '1em',
            display: 'flex',
            justifyContent: 'space-between', alignItems: 'center'}}>
            <Typography level="h2">Bill & Chill</Typography>
            <IconButton>
                <PersonOutlineIcon />
            </IconButton>
        </Box>
    );
}