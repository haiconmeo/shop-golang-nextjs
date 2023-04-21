import { TagLabel, Flex,Heading,Tag,Text,Button } from '@chakra-ui/react'
import { ArrowUpIcon,ChatIcon} from '@chakra-ui/icons'
export default function ListItem({ item }) {
    return (
      <Flex direction="row" align={"center"} mb={4}>
        <Flex style={{ flex: 2 }} justify="center" mt={-8}>
          <Tag
            size={"md"}
            key={"md"}
            borderRadius='full'
           variant='solid'
            colorScheme='teal'
          >
            <TagLabel>{item.index}</TagLabel>
          </Tag>
        </Flex>
        <div style={{ flex: 12 }}>
          <Flex direction={"column"}>
            <Heading as='h1' size='sm'>
              {item.heading}
            </Heading>
            <Flex direction={"row"} justify="space-between" mt="2" wrap={"wrap"}>
              <Text fontSize='sm' color="gray.500" >{item.site}</Text>
              <Text fontSize='sm'>{item.time} - by <span style={{ color: '#2b6cb0' }}>{item.user}</span> </Text>
            </Flex>
            <Flex direction="row">
              <Button leftIcon={<ArrowUpIcon />} colorScheme='blue' variant='ghost'>
                {item.likes}
              </Button>
           <Button leftIcon={<ChatIcon />} colorScheme='orange' variant='ghost'>
                {item.comments}
              </Button>
            </Flex>
          </Flex>
        </div>
      </Flex>
    )
  }