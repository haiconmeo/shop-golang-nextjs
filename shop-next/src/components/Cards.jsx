import {
    Box,
    Heading,
    Text,
    Stack,
    Image,
  } from '@chakra-ui/react';
  
  const IMAGE =
    'https://images.unsplash.com/photo-1518051870910-a46e30d9db16?ixlib=rb-1.2.1&ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&auto=format&fit=crop&w=1350&q=80';
  
  export default function ProductSimple({ imageSrc, imageAlt, title, category, price }) {
    return (
      <Stack p={{ base: "0 2rem" }}>
      <Image objectFit="cover" src={imageSrc} alt={imageAlt} />
      <Text color="teal.600" textTransform="uppercase">
        {category}
      </Text>
  
      <Heading color="teal.300" size="md" textTransform="capitalize">
        {title}
      </Heading>
      <Box>
        {price}
        <Box as="span" color="gray.600" fontSize="sm">
          / unit
        </Box>
      </Box>
    </Stack>
    );
  }